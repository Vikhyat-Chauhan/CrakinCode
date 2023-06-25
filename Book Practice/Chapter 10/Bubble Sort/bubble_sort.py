class Chapter10:
  def __init__(self):
    print("Program Started")
    
  def bubbleSort(self, nums1):
    print(len(nums1))
    lengthArray = len(nums1)
    
    for i in range(lengthArray):
      for j in range(lengthArray-1):
        if(nums1[j]>nums1[j+1]):            #if first element of the array is greater than the second
          temp = nums1[j+1]
          nums1[j+1] = nums1[j]
          nums1[j] = temp
    print(nums1)

p1 = Chapter10()
p1.bubbleSort([2,3,1,4,7,5,6])