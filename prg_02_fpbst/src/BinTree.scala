import scala.collection.mutable.ArrayBuffer
import scala.{+:, ::}

/*
 * CS3210 - Principles of Programming Languages - Fall 2021
 * Instructor: Thyago Mota
 * Description: Prg 02 - Functional Programming BST
 * Student(s) Name(s):
 */

class BinTree[T <: Comparable[T]]() {

  private var value: T = _
  private var left:  BinTree[T] = _
  private var right: BinTree[T] = _

  def this(value: T) {
    this();
    this.value = value
  }

  def add(value: T): Unit = {

    def add(current: BinTree[T]): BinTree[T] = {
      if (current == null)
        new BinTree[T](value)
      else {
        if (value.compareTo(current.value) < 0)
          current.left = add(current.left)
        else if (value.compareTo(current.value) > 0)
          current.right = add(current.right)
        current
      }
    }

    if (this.value == null)
      this.value = value
    else
      add(this)
  }

  // TODO #1: returns a string representation of the tree using tabs to show its structure
  def mkString(): String = {


  }

  // TODO #2: returns a new tree (from the callee) with only the values that passed the test denoted by function f
  def filter(f: T => Boolean) = {


  }

  // TODO #3: returns a new tree (from the callee) by applying function f to each of the callee's elements
  def map(f: T => T): BinTree[T] = {


  }

  // TODO #4: applies function f to each of the tree's elements
  def foreach(f: T => T): Unit = {


  }

  // TODO #5: similar to foldLeft for collections
  def foldLeft(value: T)(f: (T, T) => T ): T = {


  }

  // TODO #6: similar to foldRight for collections
  def foldRight(value: T)(f: (T, T) => T ): T = {

  }

  // TODO #7: similar to foldLeft
  def fold(value: T)(f: (T, T) => T ): T = {
  }

  // TODO #8: returns the height of the tree
  def height(): Int = {


  }

  // TODO #9: returns the number of elements of the tree
  def size(): Int = {


  }

  // TODO #10: returns true/false depending whether value is/is not found in the tree
  def search(value: T): Boolean = {

  }
}

object BinTree {

  def main(args: Array[String]): Unit = {

    // tree 1 (a BST of integers)
    var tree1 = new BinTree[Integer]()

    // TODO #11a: use array's foreach to add the following elements: 2, 1, 7, 3, 9, 10; display the tree next

    // TODO #11b: use foreach to multiply all of its elements by 2; display the tree next

    // TODO #11c: use filter to obtain another tree whose elements are > 5; display the resulting tree next

    // TODO #11d: use map to subtract 2 from each element of the original tree (modified by TODO #11b); display the resulting tree next

    // TODO #11e: use fold to compute the sum of all of the elements of the (original) tree (modified by TODO #11d)

    // TODO #11f: use fold and size to compute the average of the elements of the (original) tree (modified by TODO #11d)

    // TODO #11g: show the height of the tree

    // TODO #11h: search for elements 4 and 17 in the (original) tree (modified by TODO #11d)

    // tree 2 (a BST of strings)
    var tree2 = new BinTree[String]()

    // TODO #12a: use array's foreach to add the following elements: "perry rhodan", "icho tolot", "denetree", "deshan apian", "levian paronn", "roder roderich", "gressko gurrad"; display the tree next

    // TODO #12b: use foreach to capitalize the first letter of each word of each tree node; display the tree next

    // TODO #12c: use filter to obtain another tree whose names are made of at least 2 words; display the resulting tree next

    // TODO #12d: use map to capitalize all of the letters from each of the names in the (original) tree; display the resulting tree next

    // TODO #12e: use fold to compute the sum of all of the letters from each of the names in the tree

    // TODO #12f: use fold and size to compute the average number of characters of the names in the tree

    // TODO #12g: show the height of the tree

    // TODO #12h: search for elements "PERRY RHODAN" and "ATLAN"

  }
}
