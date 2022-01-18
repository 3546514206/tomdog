package edu.zjnu;

import edu.zjnu.core.BootStrap;
import edu.zjnu.core.CommonYesOrNo;

public class Tomdog {

	public static void main(String[] args) {
		//System.out.println(null instanceof Object);
		new BootStrap().load();

		System.out.println(CommonYesOrNo.NO.getCode());
		System.out.println(CommonYesOrNo.YES.getDesc());
		System.out.println(CommonYesOrNo.getEnumByCode(1));




	}

}
