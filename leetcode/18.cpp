class Solution {
public:
    bool canJump(vector<int>& nums) {
     int l=0;
     for(int i=0;i<nums.size();++i)
     {
        if(i>l) return false;
        l=max(l,i+nums[i]);
     }
      return true;
    }
};
