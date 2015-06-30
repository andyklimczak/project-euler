import java.util.ArrayList;
import java.util.Iterator;

public class eul7
{
	public static void main(String[] args)
	{

		int ultimateprime = 0;
		ArrayList<Integer> primelist = new ArrayList<Integer>();

		for (int i = 2; i < Integer.MAX_VALUE && primelist.size() < 10001; i++)
		{
			boolean primetest = true;
			System.out.printf("Testing %d\n", i);
			if (primelist.size() == 0)
			{
				primelist.add(i);
			}
			else
			{
				for (Iterator<Integer> iter = primelist.iterator(); iter.hasNext();)
				{
					int prime = iter.next();
					if (i % prime == 0)
					{
						System.out.printf("Testing %d mod %d \n", i, prime);
						primetest = false;
					}
				}
				if (primetest == true)
				{
					System.out.printf("Adding %d to the list of primes\n", i);
					primelist.add(i);
				}
			}
			System.out.printf("There are %d prime numbers in the list\n", primelist.size());
			ultimateprime = i;
		}
		System.out.printf("The 10001st prime number is %d\n", ultimateprime);
	}

}
