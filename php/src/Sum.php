<?php

declare(strict_types=1);

namespace Lleitep3\Study;

class Sum
{

  static public function sum()
  {
    $numbers = func_get_args();

    return array_sum($numbers);
  }
}
