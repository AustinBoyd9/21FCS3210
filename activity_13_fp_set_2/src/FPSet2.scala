object FPSet2 {

  def isPrime(n: Int): Boolean = {
    n match {
      case x if x < 2 => false
      case _ =>
        !(2 to math.sqrt(n).toInt).exists(n % _ == 0)  // kudos Malcolm!
    }
  }

  def gcd(a: Int): Int => Int = (b: Int) => {
    if (a % b == 0)
      b
    else
      gcd(b)(a % b)
  }

  def coprime(a: Int)(b: Int) = gcd(a)(b) == 1

  def totientPhi(m: Int) = (1 until m).filter(coprime(m)(_))

  def primeFactors(n: Int) = (2 to n).filter(isPrime(_)).filter(n % _ == 0)

  def ndiv(a: Int): Int => Int = (b: Int) => {
    a % b match {
      case 0 => 1 + ndiv(a / b)(b)
      case _ => 0
    }
  }

  def primeFactorsMult(n: Int) = primeFactors(n).map((x: Int) => (x, ndiv(n)(x)))

  def primesRange(a: Int) = (b: Int) => (a to b).filter(isPrime(_))

  def goldbach(n: Int) = {
    n match {
      case x if x % 2 == 0 && x > 2 => primesRange(2)(n).map((x:Int) => {
        primesRange(2)(n).map((y: Int) => (x, y)).toSeq.filter((z: (Int, Int)) => z._1 + z._2 == n)
      }).toSeq.filter(_.length > 0)
      case _ => IndexedSeq()
    }
  }

  def goldbachList(a: Int) = (b: Int) => (a to b).map((n: Int) => (n, goldbach(n))).filter(_._2.length > 0)

  def main(args: Array[String]): Unit = {
//    println(isPrime(13))
//    println(gcd(1071)(462))
//    println(totientPhi(325))
//    println(primeFactors(40))
    //println(primeFactorsMult(40))
    println(goldbachList(10)(28))
  }

}
