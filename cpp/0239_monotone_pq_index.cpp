#include <bits/stdc++.h>
using namespace std;

class Solution {  // from huahua
 public:
  vector<int> maxSlidingWindow(vector<int>& nums, int k) {
    auto q = deque<int>{};  // index queue
    auto ans = vector<int>{};
    for (auto i = 0ll; i < nums.size(); ++i) {
      while (!q.empty() && nums[q.back()] <= nums[i]) q.pop_back();  // <= !
      q.push_back(i);

      auto start = i - k + 1;  // window-start
      if (start >= 0) ans.push_back(nums[q.front()]);
      if (start == q.front()) q.pop_front();
    }
    return ans;
  }
};
