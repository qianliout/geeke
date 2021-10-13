class Solution {
    public int uniquePathsWithObstacles(int[][] obstacleGrid) {

        int m = obstacleGrid.length;
        int n = obstacleGrid[0].length;
        //考虑特殊情况
        if (obstacleGrid[0][0] == 1 || obstacleGrid[m - 1][n - 1] == 1) return 0;
        //定义二维数组
        if (obstacleGrid[0][0] == 0 && m == 1 && n == 1) return 1;
        int[][] dp = new int[m][n];
        //第一行边界位置初始化
        for (int i = 0; i < m; i++) {
            if (obstacleGrid[i][0] != 1) {
                dp[i][0] = 1;
            } else {
                break;
            }
        }
        //第一列边界位置初始化
        for (int j = 0; j < n; j++) {
            if (obstacleGrid[0][j] != 1) {
                dp[0][j] = 1;
            } else {
                break;
            }
        }
        //遍历赋值
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                if (obstacleGrid[i][j] != 1) {
                    dp[i][j] = dp[i][j - 1] + dp[i - 1][j];
                }
            }
        }

        return dp[m - 1][n - 1];
    }
}


