/**
 * The base Component interface defines operations that can be altered by
 * decorators.
 */ function _assert_this_initialized(self) {
    if (self === void 0) {
        throw new ReferenceError("this hasn't been initialised - super() hasn't been called");
    }
    return self;
}
function _class_call_check(instance, Constructor) {
    if (!(instance instanceof Constructor)) {
        throw new TypeError("Cannot call a class as a function");
    }
}
function _defineProperties(target, props) {
    for(var i = 0; i < props.length; i++){
        var descriptor = props[i];
        descriptor.enumerable = descriptor.enumerable || false;
        descriptor.configurable = true;
        if ("value" in descriptor) descriptor.writable = true;
        Object.defineProperty(target, descriptor.key, descriptor);
    }
}
function _create_class(Constructor, protoProps, staticProps) {
    if (protoProps) _defineProperties(Constructor.prototype, protoProps);
    if (staticProps) _defineProperties(Constructor, staticProps);
    return Constructor;
}
function _define_property(obj, key, value) {
    if (key in obj) {
        Object.defineProperty(obj, key, {
            value: value,
            enumerable: true,
            configurable: true,
            writable: true
        });
    } else {
        obj[key] = value;
    }
    return obj;
}
function _get(target, property, receiver) {
    if (typeof Reflect !== "undefined" && Reflect.get) {
        _get = Reflect.get;
    } else {
        _get = function get(target, property, receiver) {
            var base = _super_prop_base(target, property);
            if (!base) return;
            var desc = Object.getOwnPropertyDescriptor(base, property);
            if (desc.get) {
                return desc.get.call(receiver || target);
            }
            return desc.value;
        };
    }
    return _get(target, property, receiver || target);
}
function _get_prototype_of(o) {
    _get_prototype_of = Object.setPrototypeOf ? Object.getPrototypeOf : function getPrototypeOf(o) {
        return o.__proto__ || Object.getPrototypeOf(o);
    };
    return _get_prototype_of(o);
}
function _inherits(subClass, superClass) {
    if (typeof superClass !== "function" && superClass !== null) {
        throw new TypeError("Super expression must either be null or a function");
    }
    subClass.prototype = Object.create(superClass && superClass.prototype, {
        constructor: {
            value: subClass,
            writable: true,
            configurable: true
        }
    });
    if (superClass) _set_prototype_of(subClass, superClass);
}
function _possible_constructor_return(self, call) {
    if (call && (_type_of(call) === "object" || typeof call === "function")) {
        return call;
    }
    return _assert_this_initialized(self);
}
function _set_prototype_of(o, p) {
    _set_prototype_of = Object.setPrototypeOf || function setPrototypeOf(o, p) {
        o.__proto__ = p;
        return o;
    };
    return _set_prototype_of(o, p);
}
function _super_prop_base(object, property) {
    while(!Object.prototype.hasOwnProperty.call(object, property)){
        object = _get_prototype_of(object);
        if (object === null) break;
    }
    return object;
}
function _type_of(obj) {
    "@swc/helpers - typeof";
    return obj && typeof Symbol !== "undefined" && obj.constructor === Symbol ? "symbol" : typeof obj;
}
function _is_native_reflect_construct() {
    if (typeof Reflect === "undefined" || !Reflect.construct) return false;
    if (Reflect.construct.sham) return false;
    if (typeof Proxy === "function") return true;
    try {
        Boolean.prototype.valueOf.call(Reflect.construct(Boolean, [], function() {}));
        return true;
    } catch (e) {
        return false;
    }
}
function _create_super(Derived) {
    var hasNativeReflectConstruct = _is_native_reflect_construct();
    return function _createSuperInternal() {
        var Super = _get_prototype_of(Derived), result;
        if (hasNativeReflectConstruct) {
            var NewTarget = _get_prototype_of(this).constructor;
            result = Reflect.construct(Super, arguments, NewTarget);
        } else {
            result = Super.apply(this, arguments);
        }
        return _possible_constructor_return(this, result);
    };
}
/**
 * Concrete Components provide default implementations of the operations. There
 * might be several variations of these classes.
 */ var ConcreteComponent = /*#__PURE__*/ function() {
    "use strict";
    function ConcreteComponent() {
        _class_call_check(this, ConcreteComponent);
    }
    _create_class(ConcreteComponent, [
        {
            key: "operation",
            value: function operation() {
                return "ConcreteComponent";
            }
        }
    ]);
    return ConcreteComponent;
}();
/**
 * The base Decorator class follows the same interface as the other components.
 * The primary purpose of this class is to define the wrapping interface for all
 * concrete decorators. The default implementation of the wrapping code might
 * include a field for storing a wrapped component and the means to initialize
 * it.
 */ var Decorator = /*#__PURE__*/ function() {
    "use strict";
    function Decorator(component) {
        _class_call_check(this, Decorator);
        _define_property(this, "component", void 0);
        this.component = component;
    }
    _create_class(Decorator, [
        {
            key: "operation",
            value: /**
   * The Decorator delegates all work to the wrapped component.
   */ function operation() {
                return this.component.operation();
            }
        }
    ]);
    return Decorator;
}();
/**
 * Concrete Decorators call the wrapped object and alter its result in some way.
 */ var ConcreteDecoratorA = /*#__PURE__*/ function(Decorator) {
    "use strict";
    _inherits(ConcreteDecoratorA, Decorator);
    var _super = _create_super(ConcreteDecoratorA);
    function ConcreteDecoratorA() {
        _class_call_check(this, ConcreteDecoratorA);
        return _super.apply(this, arguments);
    }
    _create_class(ConcreteDecoratorA, [
        {
            key: "operation",
            value: /**
   * Decorators may call parent implementation of the operation, instead of
   * calling the wrapped object directly. This approach simplifies extension
   * of decorator classes.
   */ function operation() {
                return "ConcreteDecoratorA(".concat(_get(_get_prototype_of(ConcreteDecoratorA.prototype), "operation", this).call(this), ")");
            }
        }
    ]);
    return ConcreteDecoratorA;
}(Decorator);
/**
 * Decorators can execute their behavior either before or after the call to a
 * wrapped object.
 */ var ConcreteDecoratorB = /*#__PURE__*/ function(Decorator) {
    "use strict";
    _inherits(ConcreteDecoratorB, Decorator);
    var _super = _create_super(ConcreteDecoratorB);
    function ConcreteDecoratorB() {
        _class_call_check(this, ConcreteDecoratorB);
        return _super.apply(this, arguments);
    }
    _create_class(ConcreteDecoratorB, [
        {
            key: "operation",
            value: function operation() {
                return "ConcreteDecoratorB(".concat(_get(_get_prototype_of(ConcreteDecoratorB.prototype), "operation", this).call(this), ")");
            }
        }
    ]);
    return ConcreteDecoratorB;
}(Decorator);
/**
 * The client code works with all objects using the Component interface. This
 * way it can stay independent of the concrete classes of components it works
 * with.
 */ function clientCode(component) {
    // ...
    console.log("RESULT: ".concat(component.operation()));
// ...
}
/**
 * This way the client code can support both simple components...
 */ var simple = new ConcreteComponent();
console.log("Client: I've got a simple component:");
clientCode(simple);
console.log("");
/**
 * ...as well as decorated ones.
 *
 * Note how decorators can wrap not only simple components but the other
 * decorators as well.
 */ var decorator1 = new ConcreteDecoratorA(simple);
var decorator2 = new ConcreteDecoratorB(decorator1);
console.log("Client: Now I've got a decorated component:");

