import csv


def get_int_list_from_csv(file):
    list = []
    with open(file, mode="r") as f:
        csv_reader = csv.reader(f)
        for row in csv_reader:
            for col in row:
                list.append(col)

    return list


def main():
    # get lists
    li1 = get_int_list_from_csv("./assets/list1.csv")
    li2 = get_int_list_from_csv("./assets/list2.csv")
    print(li1, li2)


if __name__ == "__main__":
    main()
