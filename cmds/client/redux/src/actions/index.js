"use strict";
exports.incrementCounter = function (delta) { return ({
    type: 'INCREMENT_COUNTER',
    delta: delta
}); };
exports.resetCounter = function () { return ({
    type: 'RESET_COUNTER'
}); };
