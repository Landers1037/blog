# -*- coding:utf-8 -*-
import os
import argparse

'''
A small python script to change the old_patter to the new_pattern.
The code is so simple that I don't want to give any annotation:)

created by xk-wang@2021/3/14 afternoon.
'''

old_pattern = '/static/res/blog/'
new_pattern = '/images/'

parser = argparse.ArgumentParser()
parser.add_argument('--old', type=str, help='The original blog dir path')
parser.add_argument('--new', type=str, help='The new blog dir path')


if __name__ == '__main__':
	args = parser.parse_args()
	old_dir = args.old
	new_dir = args.new
	os.makedirs(new_dir, exist_ok=True)

	for md_file in os.listdir(old_dir):
		if ".md" in md_file:
			with open(os.path.join(old_dir, md_file), encoding='utf-8') as f:
				content = f.read()
			if old_pattern in content:
				print("modify: ", md_file)	
				content = content.replace(old_pattern, new_pattern)
				with open(os.path.join(new_dir, md_file), 'wt', encoding='utf-8') as f:
					f.write(content)