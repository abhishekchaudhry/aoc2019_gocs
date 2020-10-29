using System;
using System.IO;

namespace aoc2019_cs_day1_2
{
    class Program
    {
        static void Main(string[] args)
        {
            var sr = new StreamReader("../../../../../../input/day1.txt");
            string line;
            int total_fuel = 0;
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


    
