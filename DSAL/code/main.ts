// bridge pattern

class Abstraction {
  private implementation: Implemention;
  constructor(printer: Implemention) {
    this.implementation = printer;
  }
  public print() {
    this.implementation.Print();
  }
}

interface Implemention {
  Print(): void;
}

class mac extends Abstraction {
  constructor(printer: Implemention) {
    super(printer);
  }
}

class php implements Implemention {
  public Print(): void {
    console.log("printing files from php");
  }
}
const printer = new php();
const pc = new mac(printer);

pc.print();
