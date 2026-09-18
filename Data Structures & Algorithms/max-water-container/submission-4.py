
class Solution:
    def maxArea(self, heights: List[int]) -> int:
        maxVolume = 0
        # for h, val in enumerate(heights):
        #     j = h+1
        #     # print('len(heights):', len(heights))
        #     while j < len(heights):
        #         # print('J:', j)
        #         widthDiff = j - h
        #         minHeight = min(heights[j], val)
        #         volume = widthDiff * minHeight
        #         if maxVolume < volume:
        #             maxVolume = volume
        #             # print('j:', j)
        #             # print('h:', h)
        #             # print('widthDiff:', widthDiff)
        #             # print('minHeight:', minHeight)
        #             # print('maxVolume:', maxVolume)
        #         j += 1
        left = 0
        right = len(heights) - 1
        while left < right:
            width = right - left
            minHeight = min(heights[left], heights[right])
            maxVolume = max(maxVolume, (width*minHeight))
            if heights[left] < heights[right]: # this make sure the left NEVER by pass the right idx
                left += 1
            else:
                right -= 1
                
            
        return maxVolume