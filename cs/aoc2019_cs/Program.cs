using System;
using System.IO;

namespace aoc2019_cs
{
    class Program
    {
        static void Main(string[] args)
        {
            var sr= new StreamReader("../../../../../input/day1.txt");
            string line;
            int fuel = 0;
            while ((line = sr.ReadLine()) != null)
            {
                var mass = int.Parse(line);

                fuel += (mass / 3) - 2;
            }
            Console.WriteLine(fuel);
        }
    }
}
