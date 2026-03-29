<?php

class Solution {

    /**
     * @param String $s1
     * @param String $s2
     * @return Boolean
     */
    function areAlmostEqual($s1, $s2) {
        if ($s1 === $s2) {
            return true;
        }
        if (strlen($s1) !== strlen($s2)) {
            return false;
        }
        $diff = [];
        for ($i = 0; $i < strlen($s1); $i++) {
            if ($s1[$i] !== $s2[$i]) {
                if (count($diff) >= 2) {
                    return false;
                }
                $diff[] = [$s1[$i], $s2[$i]];
            }
        }
        
        return count($diff) === 2 && $diff[0][0] === $diff[1][1] && $diff[0][1] === $diff[1][0];
    }
}