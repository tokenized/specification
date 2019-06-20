"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const yaml_1 = __importDefault(require("yaml"));
const ramda_1 = __importDefault(require("ramda"));
const resource_data_1 = require("./resource-data");
class RejectionCodeData {
}
exports.RejectionCodeData = RejectionCodeData;
class CurrencyTypeData {
}
exports.CurrencyTypeData = CurrencyTypeData;
class EntityTypeData {
}
exports.EntityTypeData = EntityTypeData;
class PolityType {
}
exports.PolityType = PolityType;
class RoleType {
}
exports.RoleType = RoleType;
class TagType {
}
let tagTypes;
let roleTypes;
let polityTypes;
let entityTypes;
let currencyTypes;
let rejectionTypes;
class Resources {
    static GetRejectionCodes() {
        if (rejectionTypes)
            return rejectionTypes;
        rejectionTypes = ramda_1.default.indexBy((v) => '' + v.code, yaml_1.default.parse(resource_data_1.Data.yamlRejections));
        return rejectionTypes;
    }
    static GetRejectionCode(code) {
        const types = Resources.GetRejectionCodes();
        return types['' + code];
    }
    static GetCurrencies() {
        if (currencyTypes)
            return currencyTypes;
        currencyTypes = ramda_1.default.indexBy((v) => v.code, yaml_1.default.parse(resource_data_1.Data.yamlCurrencies));
        return currencyTypes;
    }
    static GetCurrency(cur) {
        const types = Resources.GetCurrencies();
        return types[cur];
    }
    static GetEntityTypes() {
        if (entityTypes)
            return entityTypes;
        entityTypes = ramda_1.default.indexBy((v) => v.code.charAt(0), yaml_1.default.parse(resource_data_1.Data.yamlEntities));
        return entityTypes;
    }
    static GetEntityType(ent) {
        const types = Resources.GetEntityTypes();
        return types[String.fromCharCode(ent)];
    }
    static GetPolityTypes() {
        if (polityTypes)
            return polityTypes;
        polityTypes = ramda_1.default.indexBy((v) => v.code, yaml_1.default.parse(resource_data_1.Data.yamlPolities));
        return polityTypes;
    }
    static GetPolityType(pol) {
        const types = Resources.GetPolityTypes();
        return types[pol];
    }
    static GetRoleTypes() {
        if (roleTypes)
            return roleTypes;
        roleTypes = ramda_1.default.indexBy((v) => '' + v.code, yaml_1.default.parse(resource_data_1.Data.yamlRoles));
        return roleTypes;
    }
    static GetRoleType(role) {
        const types = Resources.GetRoleTypes();
        return types['' + role];
    }
    static GetTagTypes() {
        if (tagTypes)
            return tagTypes;
        tagTypes = ramda_1.default.indexBy((v) => '' + v.code, yaml_1.default.parse(resource_data_1.Data.yamlTags));
        return tagTypes;
    }
    static GetTagType(tag) {
        const types = Resources.GetTagTypes();
        return types['' + tag];
    }
}
exports.Resources = Resources;
