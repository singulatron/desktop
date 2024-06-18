package modeltypes

var PlatformLlamaCpp = Platform{
	ID: "llama-cpp",
	Container: Container{
		Port: 8000,
		Images: Images{
			Default: "crufter/llama-cpp-python-simple",
			Cuda:    "crufter/llama-cpp-python-cuda",
		},
	},
}

var PlatformStableDiffusion = Platform{
	ID: "stable-diffusion",
	Container: Container{
		Port: 7860,
		Images: Images{
			Default: "nicklucche/stable-diffusion",
		},
	},
}

const mistralDescription = `Mistral excels in understanding and generating human-like text, making it a versatile tool across a multitude of domains. Its proficiency extends from generating coherent and contextually relevant text passages to providing detailed answers to queries, showcasing an impressive grasp of knowledge across a wide array of subjects.
Mistral stands out for its ability to perform tasks with remarkable accuracy and fewer resources, a leap forward in making state-of-the-art AI more accessible and sustainable.
`
const codellamaDescription = `CodeLlama is a powerful AI model that specializes in generating code snippets and providing detailed explanations for programming-related queries. It is designed to assist developers in writing code, debugging, and understanding complex programming concepts.`

const llamaChatUncensoredPrompt = `### HUMAN:
{prompt}
  
### RESPONSE:
`

var Models = []Model{
	//
	// MISTRAL 7B
	//
	{
		Id: "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q2_K.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "Mistral",
		Parameters:     "7B",
		Flavour:        "Instruct",
		Version:        "v0.2",
		Quality:        "Q2_K",
		Extension:      "GGUF",
		FullName:       "Mistral 7B Instruct v0.2 Q2_K",
		Size:           3.08,
		MaxRam:         5.58,
		QuantComment:   "smallest, significant quality loss - not recommended for most purposes",
		Description:    mistralDescription,
		PromptTemplate: "[INST] {prompt} [/INST]",
	},
	{
		Id: "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q3_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "Mistral",
		Parameters:     "7B",
		Flavour:        "Instruct",
		Version:        "v0.2",
		Quality:        "Q3_K_S",
		Extension:      "GGUF",
		FullName:       "Mistral 7B Instruct v0.2 Q3_K_S",
		Size:           3.16,
		MaxRam:         5.66,
		QuantComment:   "very small, high quality loss",
		Description:    mistralDescription,
		PromptTemplate: "[INST] {prompt} [/INST]",
	},
	{
		Id: "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q3_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "Mistral",
		Parameters:     "7B",
		Flavour:        "Instruct",
		Version:        "v0.2",
		Quality:        "Q3_K_M",
		Extension:      "GGUF",
		FullName:       "Mistral 7B Instruct v0.2 Q3_K_M",
		Size:           3.52,
		MaxRam:         6.02,
		QuantComment:   "very small, high quality loss",
		Description:    mistralDescription,
		PromptTemplate: "[INST] {prompt} [/INST]",
	},
	{
		Id: "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_L.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q3_K_L.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "Mistral",
		Parameters:     "7B",
		Flavour:        "Instruct",
		Version:        "v0.2",
		Quality:        "Q3_K_L",
		Extension:      "GGUF",
		FullName:       "Mistral 7B Instruct v0.2 Q3_K_L",
		Size:           3.82,
		MaxRam:         6.32,
		QuantComment:   "small, substantial quality loss",
		Description:    mistralDescription,
		PromptTemplate: "[INST] {prompt} [/INST]",
	},
	{
		Id: "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q4_K_S.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q4_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "Mistral",
		Parameters:     "7B",
		Flavour:        "Instruct",
		Version:        "v0.2",
		Quality:        "Q4_K_S",
		Extension:      "GGUF",
		FullName:       "Mistral 7B Instruct v0.2 Q4_K_S",
		Size:           4.14,
		MaxRam:         6.64,
		QuantComment:   "small, greater quality loss",
		Description:    mistralDescription,
		PromptTemplate: "[INST] {prompt} [/INST]",
	},
	{
		Id: "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q4_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q4_K_M.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "Mistral",
		Parameters:     "7B",
		Flavour:        "Instruct",
		Version:        "v0.2",
		Quality:        "Q4_K_M",
		Extension:      "GGUF",
		FullName:       "Mistral 7B Instruct v0.2 Q4_K_M",
		Size:           4.37,
		MaxRam:         6.87,
		Description:    mistralDescription,
		QuantComment:   "medium, balanced quality - recommended",
		PromptTemplate: "[INST] {prompt} [/INST]",
	},
	{
		Id: "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q5_K_S.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q5_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "Mistral",
		Parameters:     "7B",
		Flavour:        "Instruct",
		Version:        "v0.2",
		Quality:        "Q5_K_S",
		Extension:      "GGUF",
		FullName:       "Mistral 7B Instruct v0.2 Q5_K_S",
		Size:           5,
		MaxRam:         7.5,
		Description:    mistralDescription,
		QuantComment:   "large, very low quality loss - recommended",
		PromptTemplate: "[INST] {prompt} [/INST]",
	},
	{
		Id: "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q5_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q5_K_M.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "Mistral",
		Parameters:     "7B",
		Flavour:        "Instruct",
		Version:        "v0.2",
		Quality:        "Q5_K_M",
		Extension:      "GGUF",
		FullName:       "Mistral 7B Instruct v0.2 Q5_K_M",
		Size:           5.13,
		MaxRam:         7.63,
		QuantComment:   "large, very low quality loss - recommended",
		Description:    mistralDescription,
		PromptTemplate: "[INST] {prompt} [/INST]",
	},
	{
		Id: "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q6_K.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q6_K.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "Mistral",
		Parameters:     "7B",
		Flavour:        "Instruct",
		Version:        "v0.2",
		Quality:        "Q6_K",
		Extension:      "GGUF",
		FullName:       "Mistral 7B Instruct v0.2 Q6_K",
		Size:           5.94,
		MaxRam:         8.44,
		QuantComment:   "very large, extremely low quality loss",
		Description:    mistralDescription,
		PromptTemplate: "[INST] {prompt} [/INST]",
	},
	{
		Id: "huggingface/TheBloke/mistral-7b-instruct-v0.2.Q8_0.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q8_0.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "Mistral",
		Parameters:     "7B",
		Flavour:        "Instruct",
		Version:        "v0.2",
		Quality:        "Q8_0",
		Extension:      "GGUF",
		FullName:       "Mistral 7B Instruct v0.2 Q8_0",
		Size:           7.7,
		MaxRam:         10.2,
		QuantComment:   "very large, extremely low quality loss - not recommended",
		Description:    mistralDescription,
		PromptTemplate: "[INST] {prompt} [/INST]",
	},
	//
	// CodeLLAMA 7B
	//
	{
		Id: "huggingface/TheBloke/codellama-7b.Q2_K.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q2_K.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "7B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q2_K",
		Extension:      "GGUF",
		FullName:       "CodeLlama 7B Q2_K",
		Size:           2.83,
		MaxRam:         5.33,
		QuantComment:   "smallest, significant quality loss - not recommended for most purposes",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-7b.Q3_K_S.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q3_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "7B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q3_K_S",
		Extension:      "GGUF",
		FullName:       "CodeLlama 7B Q3_K_S",
		Size:           2.95,
		MaxRam:         5.45,
		QuantComment:   "very small, high quality loss",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-7b.Q3_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q3_K_M.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "7B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q3_K_M",
		Extension:      "GGUF",
		FullName:       "CodeLlama 7B Q3_K_M",
		Size:           3.3,
		MaxRam:         5.8,
		QuantComment:   "very small, high quality loss",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-7b.Q3_K_L.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q3_K_L.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "7B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q3_K_L",
		Extension:      "GGUF",
		FullName:       "CodeLlama 7B Q3_K_L",
		Size:           3.6,
		MaxRam:         6.1,
		QuantComment:   "small, substantial quality loss",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-7b.Q4_K_S.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q4_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "7B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q4_K_S",
		Extension:      "GGUF",
		FullName:       "CodeLlama 7B Q4_K_S",
		Size:           3.86,
		MaxRam:         6.36,
		QuantComment:   "small, greater quality loss",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-7b.Q4_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q4_K_M.gguf",
		},
		Name:           "CodeLlama",
		Platform:       PlatformLlamaCpp,
		Parameters:     "7B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q4_K_M",
		Extension:      "GGUF",
		FullName:       "CodeLlama 7B Q4_K_M",
		Size:           4.08,
		MaxRam:         6.58,
		QuantComment:   "medium, balanced quality - recommended",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-7b.Q4_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q5_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "7B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q5_K_S",
		Extension:      "GGUF",
		FullName:       "CodeLlama 7B Q5_K_S",
		Size:           4.65,
		MaxRam:         7.15,
		QuantComment:   "large, low quality loss - recommended",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-7b.Q5_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q5_K_M.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "7B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q5_K_M",
		Extension:      "GGUF",
		FullName:       "CodeLlama 7B Q5_K_M",
		Size:           4.78,
		MaxRam:         7.28,
		QuantComment:   "large, very low quality loss - recommended",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-7b.Q6_K.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q6_K.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "7B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q6_K",
		Extension:      "GGUF",
		FullName:       "CodeLlama 7B Q6_K",
		Size:           5.53,
		MaxRam:         8.03,
		QuantComment:   "very large, extremely low quality loss",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-7b.Q8_0.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q8_0.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "7B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q8_0",
		Extension:      "GGUF",
		FullName:       "CodeLlama 7B Q8_0",
		Size:           7.16,
		MaxRam:         9.66,
		QuantComment:   "very large, extremely low quality loss - not recommended",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	// CodeLLama 13B
	{
		Id: "huggingface/TheBloke/codellama-13b.Q2_K.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q2_K.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "13B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q2_K",
		Extension:      "GGUF",
		FullName:       "CodeLlama 13B Q2_K",
		Size:           5.43,
		MaxRam:         7.93,
		QuantComment:   "smallest, significant quality loss - not recommended for most purposes",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-13b.Q3_K_S.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q3_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "13B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q3_K_S",
		Extension:      "GGUF",
		FullName:       "CodeLlama 13B Q3_K_S",
		Size:           5.66,
		MaxRam:         8.16,
		QuantComment:   "very small, high quality loss",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-13b.Q3_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q3_K_M.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "13B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q3_K_M",
		Extension:      "GGUF",
		FullName:       "CodeLlama 13B Q3_K_M",
		Size:           6.34,
		MaxRam:         8.84,
		QuantComment:   "very small, high quality loss",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-13b.Q3_K_L.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q3_K_L.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "13B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q3_K_L",
		Extension:      "GGUF",
		FullName:       "CodeLlama 13B Q3_K_L",
		Size:           6.93,
		MaxRam:         9.43,
		QuantComment:   "small, substantial quality loss",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-13b.Q4_K_S.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q4_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "13B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q4_K_S",
		Extension:      "GGUF",
		FullName:       "CodeLlama 13B Q4_K_S",
		Size:           7.41,
		MaxRam:         9.91,
		QuantComment:   "small, greater quality loss",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-13b.Q4_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q4_K_M.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "13B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q4_K_M",
		Extension:      "GGUF",
		FullName:       "CodeLlama 13B Q4_K_M",
		Size:           7.87,
		MaxRam:         10.37,
		QuantComment:   "medium, balanced quality - recommended",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-13b.Q5_K_S.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q5_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "13B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q5_K_S",
		Extension:      "GGUF",
		FullName:       "CodeLlama 13B Q5_K_S",
		Size:           8.97,
		MaxRam:         11.47,
		QuantComment:   "large, low quality loss - recommended",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-13b.Q5_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q5_K_M.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "13B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q5_K_M",
		Extension:      "GGUF",
		FullName:       "CodeLlama 13B Q5_K_M",
		Size:           9.23,
		MaxRam:         11.73,
		QuantComment:   "large, very low quality loss - recommended",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-13b.Q6_K.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q6_K.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "13B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q6_K",
		Extension:      "GGUF",
		FullName:       "CodeLlama 13B Q6_K",
		Size:           10.68,
		MaxRam:         13.18,
		QuantComment:   "very large, extremely low quality loss",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	{
		Id: "huggingface/TheBloke/codellama-13b.Q6_K.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q8_0.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "CodeLlama",
		Parameters:     "13B",
		Flavour:        "Code",
		Version:        "1",
		Quality:        "Q8_0",
		Extension:      "GGUF",
		FullName:       "CodeLlama 13B Q8_0",
		Size:           13.83,
		MaxRam:         16.33,
		QuantComment:   "very large, extremely low quality loss - not recommended",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
	//
	// Llama 2 7B Chat Uncensored
	//
	{
		Id: "huggingface/TheBloke/llama2_7b_chat_uncensored.Q2_K.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q2_K.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "LLaMA2",
		Parameters:     "7B",
		Flavour:        "Chat",
		Uncensored:     true,
		Version:        "2",
		Quality:        "Q2_K",
		Extension:      "GGUF",
		FullName:       "LLaMA2 7B Chat Uncensored Q2_K",
		Size:           2.83,
		MaxRam:         5.33,
		QuantComment:   "smallest, significant quality loss - not recommended for most purposes",
		Description:    "A version of LLaMA2 model tailored for uncensored chat applications, optimized for smaller size and RAM usage with significant quality loss, making it less suitable for most purposes.",
		PromptTemplate: llamaChatUncensoredPrompt,
	},
	{
		Id: "huggingface/TheBloke/llama2_7b_chat_uncensored.Q3_K_S.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q3_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "LLaMA2",
		Parameters:     "7B",
		Flavour:        "Chat",
		Uncensored:     true,
		Version:        "2",
		Quality:        "Q3_K_S",
		Extension:      "GGUF",
		FullName:       "LLaMA2 7B Chat Uncensored Q3_K_S",
		Size:           2.95,
		MaxRam:         5.45,
		QuantComment:   "very small, high quality loss",
		Description:    "A specialized version of the LLaMA2 model for chat applications with a focus on reduced size and memory requirements, featuring a high quality loss.",
		PromptTemplate: llamaChatUncensoredPrompt,
	},
	{
		Id: "huggingface/TheBloke/llama2_7b_chat_uncensored.Q3_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q3_K_M.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "LLaMA2",
		Parameters:     "7B",
		Flavour:        "Chat",
		Uncensored:     true,
		Version:        "2",
		Quality:        "Q3_K_M",
		Extension:      "GGUF",
		FullName:       "LLaMA2 7B Chat Uncensored Q3_K_M",
		Size:           3.3,
		MaxRam:         5.8,
		QuantComment:   "very small, high quality loss",
		Description:    "This iteration of the LLaMA2 model is optimized for chat purposes, balancing size and efficiency with a slight compromise on quality.",
		PromptTemplate: llamaChatUncensoredPrompt,
	},
	{
		Id: "huggingface/TheBloke/llama2_7b_chat_uncensored.Q3_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q3_K_L.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "LLaMA2",
		Parameters:     "7B",
		Flavour:        "Chat",
		Uncensored:     true,
		Version:        "2",
		Quality:        "Q3_K_L",
		Extension:      "GGUF",
		FullName:       "LLaMA2 7B Chat Uncensored Q3_K_L",
		Size:           3.6,
		MaxRam:         6.1,
		QuantComment:   "small, substantial quality loss",
		Description:    "Designed for chat applications, this LLaMA2 model version offers a compact size with manageable memory requirements at the cost of some quality loss.",
		PromptTemplate: llamaChatUncensoredPrompt,
	},
	{
		Id: "huggingface/TheBloke/llama2_7b_chat_uncensored.Q3_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q4_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "LLaMA2",
		Parameters:     "7B",
		Flavour:        "Chat",
		Uncensored:     true,
		Version:        "2",
		Quality:        "Q4_K_S",
		Extension:      "GGUF",
		FullName:       "LLaMA2 7B Chat Uncensored Q4_K_S",
		Size:           3.86,
		MaxRam:         6.36,
		QuantComment:   "small, greater quality loss",
		Description:    "A compact and efficient version of the LLaMA2 model for uncensored chat, optimized to maintain a balance between size, memory usage, and quality.",
		PromptTemplate: llamaChatUncensoredPrompt,
	},
	{
		Id: "huggingface/TheBloke/llama2_7b_chat_uncensored.Q4_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q4_K_M.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "LLaMA2",
		Parameters:     "7B",
		Flavour:        "Chat",
		Uncensored:     true,
		Version:        "2",
		Quality:        "Q4_K_M",
		Extension:      "GGUF",
		FullName:       "LLaMA2 7B Chat Uncensored Q4_K_M",
		Size:           4.08,
		MaxRam:         6.58,
		QuantComment:   "medium, balanced quality - recommended",
		Description:    "The LLaMA2 7B Chat Uncensored Q4_K_M model offers a balanced compromise between file size, RAM requirements, and quality, making it a recommended choice for chat applications.",
		PromptTemplate: llamaChatUncensoredPrompt,
	},
	{
		Id: "huggingface/TheBloke/llama2_7b_chat_uncensored.Q4_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q5_K_S.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "LLaMA2",
		Parameters:     "7B",
		Flavour:        "Chat",
		Uncensored:     true,
		Version:        "2",
		Quality:        "Q5_K_S",
		Extension:      "GGUF",
		FullName:       "LLaMA2 7B Chat Uncensored Q5_K_S",
		Size:           4.65,
		MaxRam:         7.15,
		QuantComment:   "large, low quality loss - recommended",
		Description:    "A large-scale model version of LLaMA2 for chat, the Q5_K_S variant minimizes quality loss while requiring more memory, recommended for its efficient performance.",
		PromptTemplate: llamaChatUncensoredPrompt,
	},
	{
		Id: "huggingface/TheBloke/llama2_7b_chat_uncensored.Q4_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q5_K_M.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "LLaMA2",
		Parameters:     "7B",
		Flavour:        "Chat",
		Uncensored:     true,
		Version:        "2",
		Quality:        "Q5_K_M",
		Extension:      "GGUF",
		FullName:       "LLaMA2 7B Chat Uncensored Q5_K_M",
		Size:           4.78,
		MaxRam:         7.28,
		QuantComment:   "large, very low quality loss - recommended",
		Description:    "LLaMA2 7B Chat Uncensored Q5_K_M is tailored for high-demand chat applications, offering substantial capacity with minimal compromise on quality.",
		PromptTemplate: llamaChatUncensoredPrompt,
	},
	{
		Id: "huggingface/TheBloke/llama2_7b_chat_uncensored.Q4_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q6_K.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "LLaMA2",
		Parameters:     "7B",
		Flavour:        "Chat",
		Uncensored:     true,
		Version:        "2",
		Quality:        "Q6_K",
		Extension:      "GGUF",
		FullName:       "LLaMA2 7B Chat Uncensored Q6_K",
		Size:           5.53,
		MaxRam:         8.03,
		QuantComment:   "very large, extremely low quality loss",
		Description:    "Optimized for expansive chat integrations, the Q6_K version of LLaMA2 ensures extensive capacity with remarkably low quality loss, suitable for advanced applications.",
		PromptTemplate: llamaChatUncensoredPrompt,
	},
	{
		Id: "huggingface/TheBloke/llama2_7b_chat_uncensored.Q8_0.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q8_0.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "LLaMA2",
		Parameters:     "7B",
		Flavour:        "Chat",
		Uncensored:     true,
		Version:        "2",
		Quality:        "Q8_0",
		Extension:      "GGUF",
		FullName:       "LLaMA2 7B Chat Uncensored Q8_0",
		Size:           7.16,
		MaxRam:         9.66,
		QuantComment:   "very large, extremely low quality loss - not recommended",
		Description:    "The LLaMA2 7B Chat Uncensored Q8_0 variant represents the upper echelon in terms of size and memory requirements.",
		PromptTemplate: llamaChatUncensoredPrompt,
	},
	// Llama 3 8B
	{
		Id: "huggingface/QuantFactory/Meta-Llama-3-8B-Instruct.Q3_K_M.gguf",
		Assets: map[string]string{
			"MODEL": "https://huggingface.co/QuantFactory/Meta-Llama-3-8B-Instruct-GGUF/resolve/main/Meta-Llama-3-8B-Instruct.Q3_K_M.gguf",
		},
		Platform:       PlatformLlamaCpp,
		Name:           "Llama 3",
		Parameters:     "8B",
		Flavour:        "Code",
		Version:        "3",
		Quality:        "Q3_K_M",
		Extension:      "GGUF",
		FullName:       "Llama 3 8B Q3_K_M",
		Size:           4.02,
		MaxRam:         5.33,
		QuantComment:   "smallest, significant quality loss - not recommended for most purposes",
		Description:    codellamaDescription,
		PromptTemplate: "{prompt}",
	},
}
