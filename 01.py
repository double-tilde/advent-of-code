import csv


def get_int_list_from_csv(file):
    list = []
    with open(file, mode="r") as f:
        csv_reader = csv.reader(f)
        for row in csv_reader:
            for col in row:
                list.append(col)

    return list


def is_list_only_ints(lst):
    num_lst = []
    errors = []

    for item in lst:
        try:
            num_item = int(item)
            num_lst.append(num_item)
        except ValueError:
            errors.append(ValueError)

    if errors:
        print("The following strings are not valid integers:")
        for err in errors:
            print(err)
        return []

    return num_lst


def compare_lists(lst_1, lst_2):
    dct = {}
    diff = []
    res = 0

    for idx, val in enumerate(lst_1):
        dct[idx] = val

    for idx, val in enumerate(lst_2):
        if dct[idx] == val:
            continue
        if dct[idx] > val:
            diff.append(dct[idx] - val)
        if dct[idx] < val:
            diff.append(val - dct[idx])

    for idx, val in enumerate(diff):
        res += val

    return res


def main():
    sorted_li1 = None
    sorted_li2 = None

    # get lists
    li1 = get_int_list_from_csv("./assets/list1.csv")
    li2 = get_int_list_from_csv("./assets/list2.csv")

    # check all of the items in the list are numbers
    num_li1 = is_list_only_ints(li1)
    num_li2 = is_list_only_ints(li2)

    if num_li1:
        sorted_li1 = sorted(num_li1)

    if num_li2:
        sorted_li2 = sorted(num_li2)

    res = compare_lists(sorted_li1, sorted_li2)

    print(res)


if __name__ == "__main__":
    main()
