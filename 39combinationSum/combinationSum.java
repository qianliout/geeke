import java.util.ArrayList;



List<List<Integer>> allResultList = new ArrayList<>();
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        if (target <= 0) {
            allResultList.add(new ArrayList<>());
        }
        //先排序，排序后可以加剪枝
        Arrays.sort(candidates);
        dfs(new ArrayList<>(), candidates, 0, target, 0);
        return allResultList;
    }


    public void dfs(List<Integer> list, int[] candidates, int sum, int target, int start) {
        for (int i = start; i < candidates.length; i++) {
            list.add(candidates[i]);
            if ((sum + candidates[i]) == target) {
                //此时tmpList 满足
                List<Integer> tmpList = new ArrayList<>(list);
                allResultList.add(tmpList);
            } if ((sum + candidates[i]) < target) {
                dfs(list, candidates, sum + candidates[i], target, i);
            } else {
                list.remove(list.size() - 1);
                //(sum+candidates[i]) > target,因为数组有序，后面一定不满足
                break;
            }
            list.remove(list.size() - 1);
        }
    }
