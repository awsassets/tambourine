#include<linux/module.h>

static int milkshake_init(void){
	printk(KERN_INFO "Milkshake (INIT)\n");
	return 0;
}

static void milkshake_exit(void){
	printk(KERN_INFO "Milkshake (EXIT)\n");
}

module_init(milkshake_init);
module_exit(milkshake_exit);
MODULE_LICENSE("GPL");