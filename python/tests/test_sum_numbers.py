
from app.src.sum_numbers import sum_numbers


class TestSum:
  def test_sum_one_number(self):
    assert sum_numbers(1) == 1

  def test_sum_many_numbers(self):
    assert sum_numbers(1, 2, 3) == 6

  def test_sum_negative_numbers(self):
    assert sum_numbers(-1, 2, 3) == 4