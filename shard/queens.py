# 作者：刘强
# data: 2019-11-19

import time


class Queens1(object):
    def solve_queens(self, n):
        if n <= 1:
            return []
        self.result = []
        self.cols = set()
        self.pie = set()
        self.na = set()
        self.DFS(n, 0, [])
        return self._gen_res(n)

    def DFS(self, n, row, cur_state):
        if row >= n:
            self.result.append(cur_state)
            return
        for col in range(n):
            if col in self.cols or row + col in self.pie or row - col in self.na:
                # go die
                continue
            # update the flags
            self.cols.add(col)
            self.pie.add(col + row)
            self.na.add(row - col)

            self.DFS(n, row + 1, cur_state + [col])

            self.cols.remove(col)
            self.pie.remove(row + col)
            self.na.remove(row - col)

    def _gen_res(self, n):
        board = []
        for res in self.result:
            for i in res:
                board.append("." * i + "Q" + "." * (n - i - 1))
        return [board[i:i + n] for i in range(0, len(board), n)]


class Queens2(object):
    def solve_queens(self, n):
        def DFS(queens, xy_diff, xy_sum):
            p = len(queens)
            if p == n:
                result.append(queens)
                return
            for q in range(n):
                if q not in queens and p - q not in xy_diff and p + q not in xy_sum:
                    DFS(queens + [q], xy_diff + [p - q], xy_sum + [p + q])

        result = []
        DFS([], [], [])
        return [["." * i + "Q" + "." * (n - i - 1) for i in sol] for sol in result]


class Queens3(object):
    def solve_queens(self, n):
        if n < 1:
            return []
        self.count = 0
        self.DFS(n, 0, 0, 0, 0)
        return self.count

    def DFS(self, n, row, cols, pie, na):
        if row >= n:
            self.count += 1
            return

        bits = (~(cols | pie | na)) & ((1 << n) - 1)  # 得到当前所有的空位

        while bits:
            p = bits & - bits  # 得到最低位的1
            self.DFS(n, row + 1, cols | p, (pie | p) << 1, (na | p) >> 1)
            bits = bits & (bits - 1)  # 去掉最低位的1


if __name__ == '__main__':
    q3 = Queens3()
    start = time.time()
    q3.solve_queens(14)
    end = time.time()
    print(q3.count)
    print(end - start)

    q1 = Queens1()
    start = time.time()
    q1.solve_queens(14)
    end = time.time()
    print(end - start)

    q2 = Queens2()
    start = time.time()
    q2.solve_queens(14)
    end = time.time()
    print(end - start)
