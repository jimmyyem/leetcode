<?php

$arr = [1,4,2,3,6,5];

var_dump($arr, quickSort($arr));

function quickSort($arr) {
    if(count($arr) <= 1) {
        return $arr;
    }

    $curr = $arr[0];
    $left = $right = [];
    for($i=1; $i<count($arr); $i++) {
        if($arr[$i] < $curr) {
            $left[] = $arr[$i];
        } else {
            $right[] = $arr[$i];
        }
    }

    $left =  quickSort($left);
    $right = quickSort($right);

    return array_merge($left, [$curr], $right);
}