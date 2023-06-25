"""
169. Majority Element
Easy
15.2K
455
Companies
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

 

Example 1:

Input: nums = [3,2,3]
Output: 3
Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2
 

Constraints:

n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109
 

Follow-up: Could you solve the problem in linear time and in O(1) space?
"""

#Most Effecient solution is using Bayer Moore Majority Voting Algoritm (https://www.geeksforgeeks.org/boyer-moore-majority-voting-algorithm/)

class Solution:
  def __init__(self):
    print("Program Started")
  
  def bubbleSort(self, nums1):
    lengthArray = len(nums1)
    
    for i in range(lengthArray):
      for j in range(lengthArray-1):
        if(nums1[j]>nums1[j+1]):            #if first element of the array is greater than the second
          temp = nums1[j+1]                 #swao[]
          nums1[j+1] = nums1[j]
          nums1[j] = temp
    return nums1

  def mergeSort(self,array):
    if len(array) > 1:
#[6, 5, 5, 12, 10, 9, 1]
        #  r is the point where the array is divided into two subarrays
        r = len(array)//2
        L = array[:r]
        M = array[r:]
        # Sort the two halvesclear
        self.mergeSort(L)
        self.mergeSort(M)
        i = j = k = 0

        # Until we reach either end of either L or M, pick larger among
        # elements L and M and place them in the correct position at A[p..r]
        while i < len(L) and j < len(M):
            if L[i] < M[j]:
                array[k] = L[i]
                i += 1
            else:
                array[k] = M[j]
                j += 1
            k += 1

        # When we run out of elements in either L or M,
        # pick up the remaining elements and put in A[p..r]
        while i < len(L):
            array[k] = L[i]
            i += 1
            k += 1

        while j < len(M):
            array[k] = M[j]
            j += 1
            k += 1
    return array
  #Less Efficient method does not consider the given conditions and solves regardless  
  def majorityElement2(self, nums: list[int]) -> int:
        nums = self.mergeSort(nums)
        data = [nums[0]]
        occur = [1]
        for i in range(1,len(nums)):
            if(nums[i] == data[0]):
                occur[0]+=1
            else:
                data.insert(0,nums[i])
                occur.insert(0,1)
                if(len(data)>2):
                    if(occur[2]>occur[1]):
                        data.pop(1)
                        occur.pop(1)
                    else:
                        data.pop(2)
                        occur.pop(2)    
        if(len(data)>1):                
            if(occur[1]>occur[0]):
                return data[1]
            else:
                return data[0] 
        return data[0] 
    
  # Most Effecient Solution Using Bayers Moore Algorithm (https://www.geeksforgeeks.org/boyer-moore-majority-voting-algorithm/)   
  def majorityElement(self, nums: list[int]) -> int:
      majority = counter = 0
      for i in range(len(nums)):
          if(counter == 0):
              majority = nums[i]
          if(majority != nums[i]):
              counter-=1
          else:
              counter+=1
      return majority
            

sol = Solution()
x = sol.majorityElement([0,1,0,1,0])
print(x)