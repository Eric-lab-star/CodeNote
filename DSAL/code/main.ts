// composite

abstract class Searcher {
  protected name: string;
  constructor(name: string) {
    this.name = name;
  }
  abstract search(): void;
}
class Leaf extends Searcher {
  constructor(name: string) {
    super(name);
  }
  public search(): void {
    console.log(this.name);
  }
}

class Tree extends Searcher {
  public searchers: Searcher[] = [];
  constructor(name: string) {
    super(name);
  }
  public search(): void {
    this.searchers.forEach((s) => s.search());
  }

  public add(...i: Searcher[]) {
    this.searchers = [...this.searchers, ...i];
  }
}

const file1 = new Leaf("file1");
const file2 = new Leaf("file2");

const tree = new Tree("Tree");
tree.add(file1, file2);
tree.search();
