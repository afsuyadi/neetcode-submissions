class Solution:
    def isValid(self, s: str) -> bool:
        opener = [
            "(", "[", "{"
        ]
        couple = {
            # "(": ")",
            # "{": "}",
            # "[": "]",
            "]": "[",
            ")": "(",
            "}": "{"
        }
        # if s[0] not in opener:
        #     return False
        if len(s) == 0:
            return True
        elif len(s) == 1:
            return False
        stack = []
        # maxlen = len(s)
        for val in s:
            if not stack:
                stack.append(val)
            elif stack and val in couple: # it is a closing bracker
                if stack[-1] == couple[val]:
                    stack.pop(-1)
                else:
                    return False
            elif stack and val not in couple:
                stack.append(val)
            # if stack and val == couple[stack[-1]]: # they are a couple
            #     print('val == couple[stack[-1]]=', val, couple[stack[-1]])
            #     stack.pop(-1)
            #     print(stack)
            # elif stack and val != couple[stack[-1]]: # not a couple
            #     stack.append(val)
        # print(stack)
        if stack:
            return False

        return True