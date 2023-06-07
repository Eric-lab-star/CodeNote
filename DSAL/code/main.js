// bridge pattern
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var Abstraction = /** @class */ (function () {
    function Abstraction(printer) {
        this.implementation = printer;
    }
    Abstraction.prototype.print = function () {
        this.implementation.Print();
    };
    return Abstraction;
}());
var mac = /** @class */ (function (_super) {
    __extends(mac, _super);
    function mac(printer) {
        return _super.call(this, printer) || this;
    }
    return mac;
}(Abstraction));
var php = /** @class */ (function () {
    function php() {
    }
    php.prototype.Print = function () {
        console.log("printing files from php");
    };
    return php;
}());
var printer = new php();
var pc = new mac(printer);
pc.print();
