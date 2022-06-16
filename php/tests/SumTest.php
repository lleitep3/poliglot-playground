<?php

declare(strict_types=1);

namespace Test\Lleitep3\Study;

use Lleitep3\Study\Sum;
use PHPUnit\Framework\TestCase;

use function Lleitep3\Study\sum;

final class SumTest extends TestCase
{
  public function testSumOneNumber(): void
  {
    $this->assertSame(Sum::sum(1), 1);
  }

  public function testSumMoreThenTwoNumbers(): void
  {
    $this->assertSame(Sum::sum(1, 2, 3), 6);
  }

  public function testSumNegativeNumbers(): void
  {
    $this->assertSame(Sum::sum(1, -2, 3), 2);
  }
}
