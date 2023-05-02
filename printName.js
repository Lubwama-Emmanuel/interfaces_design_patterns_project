const arr = (nums) => {
  const mySet = new Set();
  const final = [];
  for (const num of nums) {
    mySet.add(num);
  }
  console.log(mySet);

  for (const v of mySet) {
    final.push(v);
  }
  console.log(final);
};

arr([0, 0, 1, 1, 1, 2, 2, 3, 3, 4]);
