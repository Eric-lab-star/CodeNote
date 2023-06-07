// composite
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
var __spreadArray = (this && this.__spreadArray) || function (to, from, pack) {
    if (pack || arguments.length === 2) for (var i = 0, l = from.length, ar; i < l; i++) {
        if (ar || !(i in from)) {
            if (!ar) ar = Array.prototype.slice.call(from, 0, i);
            ar[i] = from[i];
        }
    }
    return to.concat(ar || Array.prototype.slice.call(from));
};
var Searcher = /** @class */ (function () {
    function Searcher(name) {
        this.name = name;
    }
    return Searcher;
}());
var Leaf = /** @class */ (function (_super) {
    __extends(Leaf, _super);
    function Leaf(name) {
        return _super.call(this, name) || this;
    }
    Leaf.prototype.search = function () {
        console.log(this.name);
    };
    return Leaf;
}(Searcher));
var Tree = /** @class */ (function (_super) {
    __extends(Tree, _super);
    function Tree(name) {
        var _this = _super.call(this, name) || this;
        _this.searchers = [];
        return _this;
    }
    Tree.prototype.search = function () {
        this.searchers.forEach(function (s) { return s.search(); });
    };
    Tree.prototype.add = function () {
        var i = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            i[_i] = arguments[_i];
        }
        this.searchers = __spreadArray(__spreadArray([], this.searchers, true), i, true);
    };
    return Tree;
}(Searcher));
var file1 = new Leaf("file1");
var file2 = new Leaf("file2");
var tree = new Tree("Tree");
tree.add(file1, file2);
tree.search();
