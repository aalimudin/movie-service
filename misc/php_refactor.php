<?php
function findFirstStringInBracket($str){
    $firstbracket = strstr($str, '(');
    $firstbracket = ltrim($firstbracket, '(');
    strstr($firstbracket, ')', true);
}