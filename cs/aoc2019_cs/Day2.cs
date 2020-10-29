using System;
using System.IO;
using System.Linq;

namespace aoc2019_cs
{
    public class Day2
    {
        public static void Run()
        {
            var sr = new StreamReader("../../../../../input/day2.txt");
            var line = sr.ReadLine();
            var program = line.Split(',');
            var intProgram = program.Select(x => int.Parse(x)).ToArray();
            //Console.WriteLine("[{0}]", string.Join(", ", intProgram));

            intProgram[1] = 12;
            intProgram[2] = 2;
            Console.WriteLine("[{0}]", string.Join(", ", intProgram));

            int pos = 0;
            var br = false;
            for (pos = 0; pos < intProgram.Length; pos += 4)
            {
                switch (intProgram[pos])
                {
                    case 1:
                        intProgram[intProgram[pos + 3]] =
                            intProgram[intProgram[pos + 1]] + intProgram[intProgram[pos + 2]];
                        continue;
                    case 2:
                        intProgram[intProgram[pos + 3]] =
                            intProgram[intProgram[pos + 1]] * intProgram[intProgram[pos + 2]];
                        continue;
                    case 99:
                        br = true;
                        break;
                    default:
                        Console.WriteLine("The starting value is incorrect");
                        break;
                }

                if (br == true)
                {
                    Console.WriteLine("the program went inside the break loop");
                }
            }

            //Console.WriteLine("the value at the starting position is:",intProgram[0]);
            Console.WriteLine("[{0}]", string.Join(", ", intProgram));
        }
    }
}