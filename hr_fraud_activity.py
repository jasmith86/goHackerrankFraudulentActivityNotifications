#!/bin/python3


# https://www.hackerrank.com/challenges/fraudulent-activity-notifications/
# Complete the activityNotifications function below.
import logging

from bisect import *


def activityNotifications(expenditure, d):
    d_is_even = d % 2 == 0
    num_alerts = 0
    mp = d // 2
    win = sorted(expenditure[: d])
    for cur_ind, cur_val in enumerate(expenditure[d:]):
        win_elem_to_del = expenditure[cur_ind]

        # incr alerts
        if d_is_even:
            if cur_val >= win[mp] + win[mp - 1]:  # cur_val >= 2*midpoint
                num_alerts += 1
        elif cur_val >= win[mp] * 2:  # odd AND cur_val >= 2*midpoint
            num_alerts += 1

        insert_idx = bisect_left(win, win_elem_to_del)
        del win[insert_idx]
        insort_left(win, cur_val)
    return num_alerts


if __name__ == '__main__':
    # case 0
    expenditure = [2, 3, 4, 2, 3, 6, 8, 4, 5]
    d = 5
    result = activityNotifications(expenditure, d)
    assert result == 2

    # case 1
    expenditure = [1, 2, 3, 4, 4]
    d = 4
    result = activityNotifications(expenditure, d)
    assert result == 0

    # case 2
    expenditure = [10, 20, 30, 40, 50]
    d = 3
    result = activityNotifications(expenditure, d)
    assert result == 1

    with open('./fraudulent-activity-notifications-testcase01.txt', mode='r') as f:
        nd = f.readline().split()
        expenditure = list(map(int, f.readline().rstrip().split()))
    n = int(nd[0])
    d = int(nd[1])
    logging.debug("n: {} {}".format(n, type(n)))
    logging.debug("d: {} {}".format(d, type(d)))
    logging.debug("exp: {}... {} len {}".format(expenditure[:5], type(expenditure[0]), len(expenditure)))
    assert n == len(expenditure)
    answer = 633
    result = activityNotifications(expenditure, d)
    assert answer == result

    print("Success!")
