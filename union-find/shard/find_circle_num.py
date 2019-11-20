import collections


class SolutionByFloodFill(object):
    def __init__(self):
        self.dx = [-1, 1, 0, 0]
        self.dy = [0, 0, -1, 1]
        self.visited = set()  # 防止重复

    def numIslands(self, grid):
        if not grid or not grid[0]: return 0
        self.max_x, self.max_y, self.grid = len(grid), len(grid[0]), grid
        return sum([self.floodfill_dfs(i, j) for i in range(self.max_x) for j in range(self.max_y)])

    def floodfill_dfs(self, x, y):
        if not self._is_valid(x, y): return 0
        self.visited.add((x, y))
        for k in range(4):
            self.floodfill_dfs(x + self.dx[k], y + self.dy[k])
        return 1

    def floodfill_bfs(self, x, y):
        if not self._is_valid(x, y): return
        self.visited.add((x, y))
        queue = collections.deque()
        queue.append((x, y))
        while queue:
            cur_x, cur_y = queue.popleft()
            for k in range(4):
                new_x, new_y = cur_x = self.dx[k], cur_y + self.dy[k]
                if self._is_valid(new_x, new_y):
                    self.visited.add((new_x, new_y))
                    queue.append((new_x, new_y))

    def _is_valid(self, x, y):
        if x < 0 or x >= self.max_x or y < 0 or y >= self.max_y:
            return False
        if self.grid[x][y] == '0' or ((x, y) in self.visited):
            return False
        return True


class UnionFind(object):
    def __init__(self, grid):
        m, n = len(grid), len(grid[0])
        self.count = 0
        self.parent = [-1] * m * n
        self.rank = [0] * m * n

        for i in range(m):
            for j in range(n):
                if grid[i][j] == '1':
                    self.parent[i * n + j] = i * n + j
                    self.count += 1

    def find(self, i):
        if self.parent[i] != i:
            self.parent[i] = self.find(self.parent[i])
        return self.parent[i]

    def union(self, x, y):
        rootX = self.find(x)
        rootY = self.find(y)

        if rootX != rootY:
            if self.rank[rootX] > self.rank[rootY]:
                self.parent[rootY] = rootX

            elif self.rank[rootX] < self.rank[rootY]:
                self.parent[rootX] = rootY
            else:
                self.parent[rootY] = rootX
                self.rank[rootX] += 1
            self.count -= 1


class SolutionByUF(object):
    def numIslands(self, grid):
        if not grid or not grid[0]:
            return 0

        uf = UnionFind(grid)
        directions = [(0, 1), (0, -1), (-1, 0), (1, 0)]
        m, n = len(grid), len(grid[0])
        for i in range(m):
            for j in range(n):
                if grid[i][j] == '0':
                    continue
                for d in directions:
                    nr, nc = i + d[0], j + d[1]
                    if m > nr >= 0 and n > nc >= 0 and grid[nr][nc] == '1':
                        uf.union(i * n + j, nr * n + nc)
        return uf.count
