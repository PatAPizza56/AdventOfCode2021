package day1

func Question1(measurements []int) (count int) {
  /*
    Question: How many measurements are larger than the previous measurement?

    Task: Calculate number of measurements larger than previous

    Example Input:
    199 //N/A - no previous measurement
    200 //increased
    208 //increased
    210 //increased
    200 //decreased
    207 //increased
    240 //increased
    269 //increased
    260 //decreased
    263 //increased

    Example Output:
    7
    
    Time-Complexity: O(n)
  */

  for i := 1; i < len(measurements); i++ {
    if measurements[i] > measurements[i-1] {
      count++;
    }
  }

  return
}

func Question2(measurements []int) (count int) {
  /*
    Question: How many sums are larger than the previous sum?

    Task: Calculate number of windows (of 3) larger than previous
    
    Example Input: 
    199  //A      
    200  //A B    
    208  //A B C  
    210  //  B C D
    200  //E   C D
    207  //E F   D
    240  //E F G  
    269  //  F G H
    260  //    G H
    263  //      H

    Example Output:
    5

    Time-Complexity: O(n)
  */

  for i := 1; i < len(measurements) - 2; i++ {
    curWindow := measurements[i] + measurements[i+1] + measurements[i+2]
    lasWindow := measurements[i] + measurements[i+1] + measurements[i-1]
    if curWindow > lasWindow {
      count++;
    }
  }

  return
}
