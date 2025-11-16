#include "configParser.h"
#include "utility.h"
#include "stdio.h"
#include "string.h"

void trimspace(char* str){
	char* end = str + strlen(str) - 1;
	while(end > str && *end == ' ') {
		*end = '\0';
		end--;
	}
}

void init_ini(const char* file, IniFile* ini)
{
    char buffer[512];
    FILE* fp = fopen(file, "r");
    if (!fp) {
        Error("Error: can't open file\n");
        return;
    }

    while (fgets(buffer, sizeof(buffer), fp)) {
        buffer[strcspn(buffer, "\n")] = 0; 
		if(buffer[0] == '\n') continue;
		
		if(buffer[0] == '('){
			char* end = strchr(buffer, ')');
			if(end){
				*end = '\0';
				ConfigEntry* entry = &ini->sections[ini->count];
				snprintf(entry->name , MAX_CHAR , "%s" , buffer+1);
				entry->size = 0; 
				ini->count++;
				continue;
			}
			continue;
		}
		if (ini->count == 0) continue;
		char* Key = strtok(buffer, ":");
		char* Value = strtok(NULL, "");
		
		ConfigEntry* current = &ini->sections[ini->count - 1];

		if (Key && Value) {
			
			while(*Key == ' ') Key++;  
			while(*Value == ' ') Value++;
			trimspace(Value);
			trimspace(Key);

            snprintf(current->parms[current->size].key , MAX_CHAR , "%s" , Key);
            snprintf(current->parms[current->size].value , MAX_CHAR , "%s" , Value);
            current->size++;
        }
    }
	

    fclose(fp);
}

ConfigEntry getElementComponents( const char* comp, IniFile* ini)
{
	ConfigEntry _local = {0};
	for(int i=0; i<ini->count; i++)
	{
		if(compare(comp , ini->sections[i].name)){
			strcpy(_local.name , ini->sections[i].name);
			_local.size = ini->sections[i].size;

			for(int j=0; j<ini->sections[i].size; j++){
				strcpy(_local.parms[j].key, ini->sections[i].parms[j].key);
    			strcpy(_local.parms[j].value, ini->sections[i].parms[j].value);
			}
			
			return _local;
		}
	}

	return _local;
}
