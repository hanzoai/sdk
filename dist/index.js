"use strict";
/**
 * Hanzo AI SDK - Unified TypeScript implementation
 */
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __exportStar = (this && this.__exportStar) || function(m, exports) {
    for (var p in m) if (p !== "default" && !Object.prototype.hasOwnProperty.call(exports, p)) __createBinding(exports, m, p);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getRustBindings = exports.useRust = void 0;
__exportStar(require("./client"), exports);
__exportStar(require("./agent"), exports);
__exportStar(require("./mcp"), exports);
__exportStar(require("./node"), exports);
__exportStar(require("./cli"), exports);
// Try to load Rust bindings if available
let rustBindings = null;
try {
    rustBindings = require('../../rs/target/release/hanzo.node');
}
catch (e) {
    // Rust bindings not available, fall back to JS implementation
    console.debug('Using JavaScript implementation (Rust bindings not found)');
}
const useRust = () => rustBindings !== null;
exports.useRust = useRust;
const getRustBindings = () => rustBindings;
exports.getRustBindings = getRustBindings;
//# sourceMappingURL=index.js.map