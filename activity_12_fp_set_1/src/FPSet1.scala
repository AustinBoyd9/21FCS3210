object FPSet1 {

  def values(low: Int)(high: Int)(fun: Int => Int) = {
    for (i <- low to high)
      yield (i, fun(i))
  }

  def factorial(n: Int) = {
    (n to 2 by -1).foldLeft(1)(_ * _)
  }

  def main(args: Array[String]): Unit = {
    println(values(-2)(3)((a: Int) => a * a).mkString(","))

    val array = Array(3, 7, 1, 2, 9, 5, 4, 3)
    println(array.reduceRight(Math.max(_, _)))

    println(factorial(5))
  }

}
