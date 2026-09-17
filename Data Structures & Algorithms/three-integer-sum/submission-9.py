from typing import List

class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        # sort the nums 
        nums.sort()
        # i = 0       
        # j = i + 1
        # k = len(nums-1)
        result = []
        # iterate through nums with i
        for i, num in enumerate(nums):
            # subArray = []
            target = - num           
        # use j (lowPoint) and k (highPoint)
            j = i + 1
            k = len(nums) - 1
            if i > 0 and nums[i] == nums[i-1]:
                continue
            while j < k:
                subArray = []
                # print("Index..", i, j, k)
                # if nums[j] == nums[j-1]:
                #     print('nums[j] == nums[j-1]', nums[j], nums[j-1])
                #     j += 1
                #     continue
                if nums[j] + nums[k] == target:
                    print("Checking vals..", target, nums[j], nums[k])
                    subArray.append(nums[i])
                    subArray.append(nums[j])
                    subArray.append(nums[k])
                    result.append(subArray)
                    j += 1
                    k -= 1
                    while j < k and nums[j] == nums[j-1]:
                        j += 1
                    while j < k and k < len(nums) and nums[k] == nums[k+1]:
                        k -= 1
                elif nums[j] + nums[k] < target:
                    j += 1
                else:
                    k -= 1
                    
                # if subArray:
                #     result.append(subArray)
            # i += 1
        return result
                