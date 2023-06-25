''' 
Solutions : 
1. Using a record of occurances(maybe a tuple/map), the entire program can be cleaned in one rotation.
   This saves the recursion to O(n), but will require a memory location on O(n). But does this qualify
   as an in place solution?
2. Using nested loops to loop through the entire program for each occurance removal, this will decrease
   the performance to O(n2) but will only consume O(1) memory location on each iteration.
'''

class Solution:
    def removeDuplicates_1(self, nums: list[int]) -> int:
        print(nums)
        occur = {}
        i=0
        while(i<len(nums)):
            if(nums[i] in occur):                   #entry already exists
                occur[nums[i]]=1+occur[nums[i]]     #increment the entry
                if(occur[nums[i]]>1):
                    nums.pop(i)                         #remove the reoccuraing value
                    i=i-1
            else:                                   #value does not really exist already in the map
                occur[nums[i]]=0                    #added an occurance entry and not removing it since single entry is required
            i=i+1
        print(nums)
            
sol = Solution()
sol.removeDuplicates_1([1,1,1,2,2,3])