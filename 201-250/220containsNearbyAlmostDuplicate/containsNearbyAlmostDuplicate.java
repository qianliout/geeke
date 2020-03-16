class Solution {

    public boolean containsNearbyAlmostDuplicate(int[] nums, int k, int t) {
        if (t < 0) {
            return false;
        }
        HashMap<Long, Long> map = new HashMap<>();
        int n = nums.length;
        long w = t + 1; // 一个桶里边数字范围的个数是 t + 1
        for (int i = 0; i < n; i++) {
            //删除窗口中第一个数字
            if (i > k) {
                map.remove(getId(nums[i - k - 1], w));
            }
            //得到当前数的桶编号
            long id = getId(nums[i], w);
            if (map.containsKey(id)) {
                return true;
            }
            if (map.containsKey(id + 1) && map.get(id + 1) - nums[i] < w) {
                return true;
            }

            if (map.containsKey(id - 1) && nums[i] - map.get(id - 1) < w) {
                return true;
            }
            map.put(id, (long) nums[i]);
        }
        TreeMap<Integer, Integer> minTreeMap = new TreeMap<>();
        return false;
    }

    private long getId(long num, long w) {
        if (num >= 0) {
            return num / w;
        } else {
            return (num + 1) / w - 1;
        }
    }
}

