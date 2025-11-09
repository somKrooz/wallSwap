#include "configParser.h"
#include "utility.h"
#include "stdio.h"
#include "string.h"


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

		if(buffer[0] == '('){
			char* end = strchr(buffer, ')');
			int len = end - buffer - 1;
			strncpy(ini->sections[ini->count].name, buffer + 1, len);
			ini->sections[ini->count].name[len] = '\0';
			ini->sections[ini->count].size = 0; 
			ini->count++;
			continue;
		}
		if (ini->count == 0) continue;
		ConfigEntry* current = &ini->sections[ini->count - 1];

		char* Key = strtok(buffer, ":");
		char* Value = strtok(NULL, "");
		if (Key && Value && current->size < MAX_KEYS) {
			
			while(*Key == ' ') Key++;  
			while(*Value == ' ') Value++;
			if(*Value == '"') Value++;

            strncpy(current->parms[current->size].key, Key, 63);
            current->parms[current->size].key[63] = 0;

            strncpy(current->parms[current->size].value, Value, 255);
            current->parms[current->size].value[255] = 0;

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
