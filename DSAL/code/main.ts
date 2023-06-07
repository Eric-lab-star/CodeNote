// adapter

//target
class Target {
  public request() {
    console.log("printing target");
  }
}

const target = new Target();

class Adapter extends Target {
  private adaptee: Adaptee;
  constructor(adaptee: Adaptee) {
    super();
    this.adaptee = adaptee;
  }
  public request() {
    this.adaptee.AdapteeRequest();
  }
}

class Adaptee {
  constructor() {}
  public AdapteeRequest() {
    console.log("pirint Adaptee");
  }
}

function Client(target: Target) {
  target.request();
}

const adaptee = new Adaptee();
const adapter = new Adapter(adaptee);

Client(adapter);

Client(target);
