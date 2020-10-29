using System;
using System.IO;

namespace aoc2019_cs
{
    public class Day1
    {
        public static void Run()
        {
            var sr = new StreamReader("../../../../../input/day1.txt");
            string line;
            var total_fuel = 0;
            while ((line = sr.ReadLine()) != null)
            {
                var mass = int.Parse(line);
                while ((mass / 3) - 2 > 0)
                {
                    mass = (mass / 3) - 2;
                    total_fuel += mass;
                }
                total_fuel += (mass / 3) - 2;
            }

            Console.WriteLine(total_fuel);
        } 
    }
}