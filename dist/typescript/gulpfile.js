const gulp = require('gulp');
const tsc = require('gulp-typescript');
const tslint = require('gulp-tslint');
const JSON_FILES = ['src/*.json', 'src/**/*.json'];
const Config = require('./gulpfile.config');
const sourcemaps = require('gulp-sourcemaps');

// pull in the project TypeScript config
const tsProject = tsc.createProject('tsconfig.json');
var config = new Config();

gulp.task('scripts', () => {
	const tsResult = tsProject.src().pipe(tsProject());
	return tsResult.js.pipe(gulp.dest('dist'));
});

gulp.task('watch', gulp.series('scripts', () => {
	gulp.watch('src/**/*.ts', gulp.series(['tslint', 'scripts']));
}));

gulp.task('instruments', function() {
	return gulp.src(JSON_FILES)
		.pipe(gulp.dest('dist'));
});
gulp.task("tslint", () =>
	gulp.src('src/**/*.ts')
	// gulp.src("source.ts")
	.pipe(tslint({
		formatter: "verbose"
	}))
	.pipe(tslint.report())
);
gulp.task('compile-ts', function() {

	var tsResult = gulp.src('src/**/*.ts')
		.pipe(sourcemaps.init())
		.pipe(tsProject());

	tsResult.dts.pipe(gulp.dest('dist'));
	return tsResult.js
		.pipe(sourcemaps.write('.'))
		.pipe(gulp.dest('dist'));
});

gulp.task('watch', gulp.series(['watch', 'instruments', 'tslint', 'scripts']));
gulp.task('default', gulp.series(['tslint', 'scripts']));
