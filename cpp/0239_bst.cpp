// 2022.01.20

#include <bits/stdc++.h>
using namespace std;

class Solution {
 public:
  vector<int> maxSlidingWindow(vector<int>& nums, int k) {
    auto window = multiset<int>(nums.begin(), nums.begin() + k);
    auto ans = vector<int>{*window.rbegin()};
    for (size_t i = k; i < nums.size(); i++) {
      auto lb = window.equal_range(nums[i - k]).first;  // lower bound
      window.erase(lb);
      window.insert(nums[i]);
      ans.push_back(*window.rbegin());
    }
    return ans;
  }
};
