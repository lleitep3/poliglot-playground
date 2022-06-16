import test from 'node:test'
import assert from 'node:assert/strict'
import { sum } from '../src/sum.js'

test('sum', async (t) => {
  await t.test('sum one number', () => {
    assert.strictEqual(sum(2), 2)
  })

  await t.test('sum two numbers', () => {
    assert.strictEqual(sum(2, 3), 5)
  })

  await t.test('sum zero', () => {
    assert.strictEqual(sum(0), 0)
  })

  await t.test('sum negative numbers', () => {
    assert.strictEqual(sum(0, -2, 3, 5, 2), 8)
  })
})