<?php
//设计一个类，只能生成该类的一个实例
class single
{
	//本类的静态实例
	private static $instance = null;

	//不能直接实例化
	private function __construct()	{}

	//获取本类实例
	public static function getInstance()
	{
		if(!self::$instance instanceof self)
		{
			self::$instance = new self();
		}
		return self::$instance;
	}
}

//测试
$single = single::getInstance();
print_r($single);