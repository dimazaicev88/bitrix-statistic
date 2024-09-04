<?php

/**
 * Comparison Operators
 */
enum Operator: string
{
    case Eq = '=';
    case NotEq = '<>';
    case Gt = '>';
    case Gte = '>=';
    case Lt = '<';
    case Lte = '=<';
    case Like = 'like';
    case NotLike = 'not like';
    case And = 'and';
    case Or = 'or';
}