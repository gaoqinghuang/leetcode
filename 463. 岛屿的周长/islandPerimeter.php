<?php
header("Content-type:text/html;charset=utf-8");

/*  给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。
网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。
*/
//时间复杂度O(n*m),没找到算法规律,偷瞄了一眼别人写的
function islandPerimeter($grid)
{
    //方法，遍历，遇到1*4，遇到相连-1
    $L = 0;
    foreach ($grid as $key => $value) {
        foreach ($value as $k => $v) {
            if ($v == 1) {
                $L += 4;
                //上
                if (isset($grid[$key - 1][$k]) && $grid[$key - 1][$k] == 1) {
                    $L -= 1;
                }
                //下
                if (isset($grid[$key + 1][$k]) && $grid[$key + 1][$k] == 1) {
                    $L -= 1;
                }
                //左
                if (isset($grid[$key][$k - 1]) && $grid[$key][$k - 1] == 1) {
                    $L -= 1;
                }
                //右
                if (isset($grid[$key][$k + 1]) && $grid[$key][$k + 1] == 1) {
                    $L -= 1;
                }
            }
        }
    }
    return $L;
}

$grid = [
    [0, 1, 0, 0],
    [1, 1, 1, 0],
    [0, 1, 0, 0],
    [0, 1, 0, 0]
];

$L = islandPerimeter($grid);
if ($L == 14) {
    echo '成功';
} else {
    echo '失败';
}
?>