import java.math.BigInteger;


public class eul16
{
	public static void main(String[] args)
	{
		BigInteger x = new BigInteger("2");
		System.out.println(x);
		x = x.pow(1000);
		System.out.println(x);
		String xstr = x.toString();
		System.out.println(xstr);
		int sum = 0;
		for(int i = 0; i < xstr.length(); i++)
		{
			char c = xstr.charAt(i);
			System.out.println(c + "   " + Character.getNumericValue(c));
			sum += Character.getNumericValue(c);
		}
		System.out.println(sum);
	}

}
