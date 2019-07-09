import _ from '@keyring/util';

class BitWriter {
	a = new _.Writer();
	pending = 0;
	position = 0;

	write(bits: boolean[]) {
		bits.forEach(b => {
			if (b) this.pending |= (1<<(7-this.position%8));
			this.position++;
			if (!(this.position%8)) {
				this.a.uint8(this.pending);
				this.pending = 0;
			}
		});
	}
	flush() {
		if (this.position%8) {
			this.a.uint8(this.pending);
			this.pending = 0;
		}
		return this;
	}
	get buf() {
		return this.a.buf;
	}
}
class BitReader {
	a: _.Reader;
	pending = 0;
	position = 0;

	constructor(buf: Buffer) {
		this.a = new _.Reader(buf);
	}
	read() {
		if (!(this.position%8)) {
			this.pending = this.a.uint8();
		}
		const bool = !!(this.pending & (1<<(7-(this.position%8))));
		this.position++;
		return bool;
	}
	get eof() {
		return this.a.eof;
	}
}

// Permission represents the permissions assigned to a specific field.
export class Permission {
	permitted: boolean;
	administration_proposal: boolean;
	holder_proposal: boolean;
	voting_systems_allowed: boolean[];

	// ReadAuthFlags reads raw auth flag data into an array of permission structs.
	static ReadAuthFlags(authFlags: number[], fields, votingSystems: number): Permission[] {
		const stream = new BitReader(Buffer.from(authFlags));
		const result = [...Array(fields)]
			.map(() => {
				const permission = new Permission();

				// Permitted
				permission.permitted = stream.read();

				// AdministrationProposal
				permission.administration_proposal = stream.read();

				// HolderProposal
				permission.holder_proposal = stream.read();

				// Voting Systems
				if (permission.administration_proposal || permission.holder_proposal) {
					permission.voting_systems_allowed = [...Array(votingSystems)]
						.map(() => stream.read());
				}

				return permission;
			});

		if (!stream.eof) {
			throw new Error('Bytes remaining');
		}

		return result;
	}

	// WriteAuthFlags writes an array of permission structs into raw auth flag data.
	static WriteAuthFlags(permissions: Permission[]): Buffer {
		const stream = new BitWriter();

		permissions.forEach(permission => {
			// Permitted
			stream.write([permission.permitted]);

			// AdministrationProposal
			stream.write([permission.administration_proposal]);

			// HolderProposal
			stream.write([permission.holder_proposal]);

			// Voting Systems
			if (permission.administration_proposal || permission.holder_proposal) {
				stream.write(permission.voting_systems_allowed);
			}
		});
		return stream.flush().buf;
	}
}

