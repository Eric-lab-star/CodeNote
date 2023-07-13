void main() {
  var p = Point(19, 10);
  print(p.x);
}

class Point {
  final double x;
  final double y;
  Point(double x, double y)
      : x = x,
        y = y;
}
