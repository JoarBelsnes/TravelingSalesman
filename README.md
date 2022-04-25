# TravelingSalesman

For this task i used the dijkstra framework to help me complete the assignment. 
I added a test file (dijkstra_salesman_test) which runs the dijkstra algorithm on the cheapest way from Timisoara to Bucharest.
I could not do benchmark tests because I only implemented one algorithm to my program, where in this task it was only required to algorithms for especially interested students.
I will try to add a second algorithm in the branch "TwoA" but it isn't sure if it going to be finished.

Test result of dijkstra algorithm of the TravelingSalesman problem
![img_2.png](img_2.png)

Update 25.04.2022:
Have researched for different types of algorithms, and how to solve it differently from dijkstra, and what i found out is that most sources use a graph method to place their "towns" and only need x and y values to place them. I havent found out how to convert the picture in the Mandatory 3 on towns in Romania and their arcs and verrticies to a x and y format. 
Examples other than dijkstra i have looked at: /n
  https://handcraftsman.wordpress.com/2012/04/02/go-implementation-of-a-travelling-salesperson-problem-solver/ /n
  https://levelup.gitconnected.com/a-nearest-neighbor-solution-in-go-to-the-traveling-salesman-problem-d4d56125b571 /n
  https://github.com/dyxj/go-tsp /n
  https://kommradhomer.medium.com/my-lazy-take-on-travelling-salesman-problem-in-golang-f7b913878c5 /n
