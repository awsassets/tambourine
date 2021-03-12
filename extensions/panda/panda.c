#include<linux/module.h>

static int nova_init(void){
	printk(KERN_INFO "Panda (INIT)\n");
	return 0;
}

static void nova_exit(void){
	printk(KERN_INFO "Panda (EXIT)\n");
}

module_init(nova_init);
module_exit(nova_exit);
MODULE_LICENSE("GPL");