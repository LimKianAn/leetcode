struct Solution;

impl Solution {
    pub fn find_circle_num(mut is_connected: Vec<Vec<i32>>) -> i32 {
        fn clear_connected(i: usize, is_connected: &mut Vec<Vec<i32>>) {
            // closure not recursive
            for j in 0..is_connected.len() {
                if is_connected[i][j] == 0 {
                    continue;
                }
                is_connected[i][j] = 0;
                is_connected[j][i] = 0;
                clear_connected(j, is_connected)
            }
        }

        let mut ans = 0;
        for i in 0..is_connected.len() {
            if is_connected[i][i] == 0 {
                continue;
            }
            ans += 1;
            clear_connected(i, &mut is_connected);
        }
        ans
    }
}
