# Go Channel Deadlock Race Condition

This repository demonstrates a potential deadlock race condition in a Go program that uses goroutines and channels.  The program appears to work correctly in some instances, but can deadlock in others due to the race condition related to channel closing and range iteration.

## Problem

The `main` function launches two goroutines. One sends integers to a channel, and the other receives them and prints them. The race condition occurs in the interaction of these goroutines with the channel.  Closing the channel in one goroutine and iterating over the channel using a range loop in the other can lead to an unpredictable outcome. Under certain conditions the receiver can end before the sender is finished. If this happens the receiving goroutine may fail to receive remaining values.

## Solution

The solution involves ensuring that the sending goroutine closes the channel after sending all values. This ensures that the receiving goroutine only waits for all sent values and will never block when the channel is closed.