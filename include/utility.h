#pragma once
#include <stdbool.h> 
#include "configParser.h"


bool checkExistance();
char* getConfigPath();
bool compare(const char* src,const char* tar);
const char* getCurrentModule(IniFile* ini);
const char* getRandomWallpaper(const char* path);
void changeWallpaper(const char* path); 
const char* getWallpaperFromWeb(const char* path);
void Log(const char* fmt, ...);
void Error(const char* fmt, ...);
