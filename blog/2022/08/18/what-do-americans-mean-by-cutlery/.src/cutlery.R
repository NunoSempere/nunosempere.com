## install.packages('ggplot2')
library(ggplot2)
library(ggthemes)

## import data
dir <- "~/Downloads/cutlery.csv"
data <- read.csv(paste(dir, "cutlery.csv", sep=""), header=TRUE)
# View(data)
answer_var = colnames(data)[16]
answer_var

freq_cut = sum(data[, answer_var] == "Knifes and cutting instruments") / dim(data)[1]
freq_eat = sum(data[, answer_var] == "Knifes, forks and spoons") / dim(data)[1]

d = list()
d$x =  c("Knifes and cutting instruments", "Knifes, forks and spoons")
d$y = c(freq_cut, freq_eat)
d2 = as.data.frame(d)

title_text = "What do Americans think that 'cutlery' refers to?"
subtitle_text = "Utensils to eat, or things which cut? (n=21)"
label_x_axis = element_blank()
label_y_axis = element_blank()
ggplot(data=d2, aes(x=x, y=y))+
  geom_col(fill="navyblue")+
  labs(
    title=title_text,
    subtitle=subtitle_text,
    x=label_x_axis, 
    y=label_y_axis 
  ) +
   scale_y_continuous(labels = scales::percent, breaks=seq(0,1,by=0.2), limits=c(0,0.6)) +
  theme_tufte() + 
  theme(
    axis.ticks.x = element_blank(),
    legend.title = element_blank(),
    plot.title = element_text(hjust = 0.5),
    plot.subtitle = element_text(hjust = 0.5), 
    legend.position="bottom",
    legend.box="vertical",
    axis.text.x=element_text(angle=0, hjust=0.5),
    plot.background=element_rect(fill = "white",colour = NA)
  )

getwd() ## Working directory on which the file will be saved. Can be changed with setwd("/your/directory")
height = 5
width = floor(height*(1+sqrt(5))/3)
ggsave("cutlery.png", width=width, height=height)
